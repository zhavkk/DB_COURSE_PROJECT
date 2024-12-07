-- Создание таблицы ролей пользователей (создаётся первой)
CREATE TABLE "user_roles" (
    "id" BIGSERIAL PRIMARY KEY,  -- Используем BIGSERIAL для автоинкремента
    "role_name" VARCHAR(255) NOT NULL  -- Роль (например, клиент, сотрудник, администратор)
);

-- Создание таблицы пользователей
CREATE TABLE "users" (
    "id" BIGSERIAL PRIMARY KEY,  -- Теперь id будет автоинкрементом
    "role_id" BIGINT NOT NULL,  -- Связь с ролью пользователя
    "login" VARCHAR(255) NOT NULL,  -- Логин
    "password_hash" VARCHAR(255) NOT NULL,  -- Хеш пароля
    FOREIGN KEY ("role_id") REFERENCES "user_roles" ("id")
);

-- Создание таблицы клиентов
CREATE TABLE "clients" (
    "id" BIGSERIAL PRIMARY KEY,  -- Используем BIGSERIAL для автоинкремента
    "name" VARCHAR(255) NOT NULL,
    "birth_date" DATE NOT NULL,
    "address" TEXT NOT NULL,
    "medical_needs" TEXT NOT NULL,
    "preferences" TEXT NOT NULL,
    "user_id" BIGINT NOT NULL,  -- Связь с таблицей users
    FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

-- Создание таблицы сотрудников
CREATE TABLE "employees" (
    "id" BIGSERIAL PRIMARY KEY,  -- Используем BIGSERIAL для автоинкремента
    "name" VARCHAR(255) NOT NULL,
    "qualification" VARCHAR(255) NOT NULL,
    "schedule" TEXT NOT NULL,
    "contact_info" TEXT NOT NULL,
    "user_id" BIGINT NOT NULL,  -- Связь с таблицей users
    FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

-- Создание таблицы услуг
CREATE TABLE "services" (
    "id" BIGSERIAL PRIMARY KEY,  -- Используем BIGSERIAL для автоинкремента
    "service_type" TEXT NOT NULL,
    "duration" INT NOT NULL  -- Тип данных для длительности
);

-- Создание таблицы заявок на услуги
CREATE TABLE "service_requests" (
    "id" BIGSERIAL PRIMARY KEY,  -- Используем BIGSERIAL для автоинкремента
    "client_id" BIGINT NOT NULL,  -- Связь с клиентом
    "service_id" BIGINT NOT NULL,  -- Связь с услугой
    "status" INT NOT NULL,  -- Статус заявки (например, в процессе или завершено)
    "request_date" DATE NOT NULL,  -- Дата подачи заявки
    "completion_date" DATE,  -- Дата завершения (может быть NULL)
    FOREIGN KEY ("client_id") REFERENCES "clients" ("id"),
    FOREIGN KEY ("service_id") REFERENCES "services" ("id")
);

-- Создание таблицы отчетов о выполнении услуги
CREATE TABLE "service_reports" (
    "id" BIGSERIAL PRIMARY KEY,  -- Используем BIGSERIAL для автоинкремента
    "request_id" BIGINT NOT NULL,  -- Связь с заявкой
    "report_text" TEXT NOT NULL,  -- Описание выполненной услуги
    "feedback" TEXT,  -- Обратная связь
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

-- Создание функции для триггера, который будет вставлять данные в таблицы clients или employees
CREATE OR REPLACE FUNCTION insert_client_or_employee() 
RETURNS TRIGGER AS
$$
BEGIN
    IF NEW.role_id = 3 THEN
        -- Вставка записи в clients с автоинкрементом id
        INSERT INTO clients (name, birth_date, address, medical_needs, preferences, user_id)
        VALUES ('Default Client Name', '1990-01-01', 'Default Address', 'No medical needs', 'No preferences', NEW.id);
    ELSIF NEW.role_id = 2 THEN
        -- Вставка записи в employees с автоинкрементом id
        INSERT INTO employees (name, qualification, schedule, contact_info, user_id)
        VALUES ('Default Employee Name', 'Employee Qualification', '9 AM - 6 PM', 'contact@company.com', NEW.id);
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Создание триггера, который срабатывает после вставки нового пользователя
CREATE TRIGGER trigger_insert_client_or_employee
AFTER INSERT ON users
FOR EACH ROW
EXECUTE FUNCTION insert_client_or_employee();
