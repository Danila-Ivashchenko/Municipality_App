package core_errors

import "net/http"

var (
	MunicipalityNotFound   = NewDomainError(http.StatusNotFound, "municipality not found", "Муниципальное образование не найдено")
	PassportNotFound       = NewDomainError(http.StatusNotFound, "passport not found", "Паспорт туризма не найден")
	ChapterNotFound        = NewDomainError(http.StatusNotFound, "chapter not found", "Глава паспорта туризма не найдена")
	PartitionNotFound      = NewDomainError(http.StatusNotFound, "partition not found", "Раздел паспорта туризма не найден")
	UserNotFound           = NewDomainError(http.StatusNotFound, "user not found", "Пользователь не найден")
	UserAuthTokenNotFound  = NewDomainError(http.StatusNotFound, "user_auth_token not found", "Сессия не найдена")
	ObjectTemplateNotFound = NewDomainError(http.StatusNotFound, "object template not found", "Шаблон объекта туризма не найден")
	EntityTemplateNotFound = NewDomainError(http.StatusNotFound, "entity template not found", "Шаблон сущности туризма не найден")

	EntityNameIsUsed = NewDomainError(http.StatusBadRequest, "entity name is used", "Имя сущности используется")
	ObjectNameIsUsed = NewDomainError(http.StatusBadRequest, "object name is used", "Имя объекта используется")

	EntityTypeNameIsUsed = NewDomainError(http.StatusBadRequest, "entity template not found", "Имя типа сущности уже использутеся")
	ObjectTypeNameIsUsed = NewDomainError(http.StatusBadRequest, "entity template not found", "Имя типа объекта уже использутеся")

	EntityTypeNotFound = NewDomainError(http.StatusNotFound, "entity template not found", "Типа сущнсоти не найден")
	ObjectTypeNotFound = NewDomainError(http.StatusNotFound, "entity template not found", "Типа объекта не найден")

	EntityTemplateNameIsUsed = NewDomainError(http.StatusBadRequest, "entity template not found", "Имя шаблон сущности уже использутеся")
	ObjectTemplateNameIsUsed = NewDomainError(http.StatusBadRequest, "entity template not found", "Имя шаблон объекта уже использутеся")

	ValidationError       = NewDomainError(http.StatusBadRequest, "validation core_errors", "Ошибка валидации")
	EmailAlreadyUsedError = NewDomainError(http.StatusBadRequest, "email already used", "Email уже используется")
	AuthErrorError        = NewDomainError(http.StatusBadRequest, "fail to authority already used", "Ошибка авторизации")

	PassportNameAlreadyUsed  = NewDomainError(http.StatusBadRequest, "passport name already used", "Данное имя паспорта туризма уже используется")
	ChapterNameAlreadyUsed   = NewDomainError(http.StatusBadRequest, "chapter name already used", "Глава паспорта туризма уже используется")
	PartitionNameAlreadyUsed = NewDomainError(http.StatusBadRequest, "partition name already used", "Раздел паспорта туризма уже используется")

	UserHasNoPermissionToWrite  = NewDomainError(http.StatusNotFound, "user has not permissions to write", "Пользователь не имеет прав на внесение изменений")
	UserHasNoPermissionToDelete = NewDomainError(http.StatusNotFound, "user has not permissions to delete", "Пользователь не имеет прав на удаление")
	UserIsNotAdmin              = NewDomainError(http.StatusNotFound, "user is not admin", "Пользователь не является администратором")
	UserNotAuth                 = NewDomainError(http.StatusNotFound, "user not auth", "Пользователь не авторизирован")
)
