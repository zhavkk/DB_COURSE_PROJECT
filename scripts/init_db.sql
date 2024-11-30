-- Создание таблицы ролей пользователей (создаётся первой)
CREATE TABLE "user_roles" (
    "id" BIGINT NOT NULL,
    "role_name" VARCHAR(255) NOT NULL,  -- Роль (например, клиент, сотрудник, администратор)
    PRIMARY KEY ("id")
);

-- Создание таблицы пользователей
CREATE TABLE "users" (
    "id" BIGINT NOT NULL,  -- ID пользователя
    "role_id" BIGINT NOT NULL,  -- Роль пользователя (например, клиент, сотрудник, администратор)
    "login" VARCHAR(255) NOT NULL,  -- Логин
    "password_hash" VARCHAR(255) NOT NULL,  -- Хеш пароля
    PRIMARY KEY ("id"),
    FOREIGN KEY ("role_id") REFERENCES "user_roles" ("id")
);

-- Создание таблицы клиентов
CREATE TABLE "clients" (
    "id" BIGINT NOT NULL,  -- ID клиента
    "name" VARCHAR(255) NOT NULL,
    "birth_date" DATE NOT NULL,
    "address" TEXT NOT NULL,
    "medical_needs" TEXT NOT NULL,
    "preferences" TEXT NOT NULL,
    "user_id" BIGINT NOT NULL,  -- Связь с таблицей users
    PRIMARY KEY ("id"),
    FOREIGN KEY ("user_id") REFERENCES "users" ("id")  -- Связь с учетной записью клиента
);

-- Создание таблицы сотрудников
CREATE TABLE "employees" (
    "id" BIGINT NOT NULL,  -- ID сотрудника
    "name" VARCHAR(255) NOT NULL,
    "qualification" VARCHAR(255) NOT NULL,
    "schedule" TEXT NOT NULL,
    "contact_info" TEXT NOT NULL,
    "user_id" BIGINT NOT NULL,  -- Связь с таблицей users
    PRIMARY KEY ("id"),
    FOREIGN KEY ("user_id") REFERENCES "users" ("id")  -- Связь с учетной записью сотрудника
);

-- Создание таблицы услуг
CREATE TABLE "services" (
    "id" BIGINT NOT NULL,
    "service_type" TEXT NOT NULL,
    "duration" INT NOT NULL,  -- Тип данных для длительности
    PRIMARY KEY ("id")
);

-- Создание таблицы заявок на услуги
CREATE TABLE "service_requests" (
    "id" BIGINT NOT NULL,
    "client_id" BIGINT NOT NULL,  -- Связь с клиентом
    "service_id" BIGINT NOT NULL,  -- Связь с услугой
    "status" INT NOT NULL,  -- Статус заявки (например, в процессе или завершено)
    "request_date" DATE NOT NULL,  -- Дата подачи заявки
    "completion_date" DATE,  -- Дата завершения (может быть NULL)
    PRIMARY KEY ("id"),
    FOREIGN KEY ("client_id") REFERENCES "clients" ("id"),
    FOREIGN KEY ("service_id") REFERENCES "services" ("id")
);

-- Создание таблицы отчетов о выполнении услуги
CREATE TABLE "service_reports" (
    "id" BIGINT NOT NULL,
    "request_id" BIGINT NOT NULL,  -- Связь с заявкой
    "report_text" TEXT NOT NULL,  -- Описание выполненной услуги
    "feedback" TEXT,  -- Обратная связь
    PRIMARY KEY ("id"),
    FOREIGN KEY ("request_id") REFERENCES "service_requests" ("id")
);

-- Промежуточная таблица для связи сотрудников и заявок
CREATE TABLE "service_request_employees" (
    "request_id" BIGINT NOT NULL,  -- Связь с заявкой
    "employee_id" BIGINT NOT NULL,  -- Связь с сотрудником
    PRIMARY KEY ("request_id", "employee_id"),
    FOREIGN KEY ("request_id") REFERENCES "service_requests" ("id"),
    FOREIGN KEY ("employee_id") REFERENCES "employees" ("id")
);

-- Ограничение на проверку даты рождения клиента
ALTER TABLE "clients" ADD CONSTRAINT "birth_date_check" CHECK ("birth_date" < CURRENT_DATE);

-- Ограничение на проверку статуса заявки (0 - в процессе, 1 - завершено)
ALTER TABLE "service_requests" ADD CONSTRAINT "status_check" CHECK ("status" IN (0, 1));
