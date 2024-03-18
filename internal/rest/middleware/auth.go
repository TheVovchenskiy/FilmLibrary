package middleware

import (
	"context"
	"database/sql"
	"filmLibrary/configs"
	"filmLibrary/model"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID int `json:"userId"`
	jwt.StandardClaims
}

func GetUserFromRequest(r *http.Request, db *sql.DB) (*model.User, error) {
	tokenString := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	fmt.Println(tokenString)

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return configs.JwtKey, nil
	})
	fmt.Println(token)
	fmt.Println(claims.UserID)

	if err != nil || !token.Valid {
		fmt.Println(err)
		return nil, fmt.Errorf("invalid token")
	}

	user, err := GetUserByID(context.Background(), claims.UserID, db)
	fmt.Println(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByID(ctx context.Context, userID int, db *sql.DB) (*model.User, error) {
	user := &model.User{}
	err := db.QueryRowContext(ctx, "SELECT id FROM users WHERE id = $1", userID).Scan(&user.Id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserRoles(userID int, db *sql.DB) ([]model.Role, error) {
	// Запрос к базе данных для получения ролей пользователя
	rows, err := db.Query("SELECT r.id, r.name FROM roles r INNER JOIN user_roles ur ON r.id = ur.role_id WHERE ur.user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []model.Role
	for rows.Next() {
		var role model.Role
		if err := rows.Scan(&role.Id, &role.Name); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func RoleCheckMiddleware(db *sql.DB) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := GetUserFromRequest(r, db)
			if err != nil {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			// Получение ролей пользователя из базы данных
			roles, err := GetUserRoles(user.Id, db)
			if err != nil {
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			// Проверка наличия роли админа
			isAdmin := false
			for _, role := range roles {
				if role.Name == "admin" {
					isAdmin = true
					break
				}
			}

			// Если пользователь не админ и метод не GET, отказываем в доступе
			if !isAdmin && r.Method != http.MethodGet {
				http.Error(w, "Forbidden", http.StatusForbidden)
				return
			}

			// Если проверка пройдена, передаем управление следующему обработчику
			next.ServeHTTP(w, r)
		})
	}
}
