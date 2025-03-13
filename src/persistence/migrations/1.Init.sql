CREATE TABLE IF NOT EXISTS "employees" (
    "id" uuid NOT NULL,
    "payroll_number" text,
    "forenames" text,
    "surname" text,    
    "date_of_birth" date NOT NULL,   
    "telephone_number" text,
    "mobile_number" text,
    "address_line_1" text,
    "address_line_2" text,
    "post_code" text,
    "email" text,
    "start_date" date NOT NULL,
    "created_at" timestamp NOT NULL,
    "updated_at" timestamp NOT NULL,
    CONSTRAINT "pk_employees" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX IF NOT EXISTS "ix_employees_payroll_number" ON "employees" ("payroll_number");