CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE "gpu_models" (
    "id" varchar(36) NOT NULL DEFAULT uuid_generate_v4(),
    "vram" float DEFAULT 6,
    "octane_bench_score" int DEFAULT 0,
    "gpu_no" int DEFAULT 0,
    "available" int DEFAULT 1,
    "vram_free" float DEFAULT NULL,
    "updated" timestamp DEFAULT current_timestamp,
    PRIMARY KEY ("id")
);

CREATE INDEX "gpu_models_vram_IDX" on "gpu_models" ("vram");

CREATE INDEX "gpu_models_vram_free_IDX" on "gpu_models" ("vram_free");

CREATE INDEX "gpu_models_updated_IDX" on "gpu_models" ("updated");

CREATE TABLE "gpu_services" (
    "id" varchar(36) NOT NULL DEFAULT uuid_generate_v4(),
    "container" varchar(100) DEFAULT NULL,
    "vram" float DEFAULT 6,
    "health" varchar(50) DEFAULT '/',
    "concurrent" boolean DEFAULT false,
    "utilization" float DEFAULT 1,
    PRIMARY KEY ("id")
);

CREATE TABLE "gpu_service_mounts" (
    "service_id" varchar(36) NOT NULL,
    "host_path" varchar(36) DEFAULT '/tmp/clust_mgr',
    "container_path" varchar(36) DEFAULT '/tmp/clust_mgr',
    PRIMARY KEY ("service_id"),
    CONSTRAINT "gpu_service_mounts_ibfk_1" FOREIGN KEY ("service_id") REFERENCES "gpu_services" ("id")
);

CREATE TABLE "gpu_service_ports" (
    "service_id" varchar(36) DEFAULT NULL,
    "host_port" int DEFAULT NULL,
    "container_port" int DEFAULT 8000,
    PRIMARY KEY ("service_id"),
    CONSTRAINT "gpu_service_ports_ibfk_1" FOREIGN KEY ("service_id") REFERENCES "gpu_services" ("id")
);

CREATE INDEX "gpu_service_ports_service_id_IDX" on "gpu_service_ports" ("service_id");

CREATE TABLE "hosts" (
    "id" char(36) NOT NULL DEFAULT uuid_generate_v4(),
    "name" varchar(100) NOT NULL,
    "ready" int DEFAULT 1,
    "updated" timestamp DEFAULT current_timestamp,
    "job_url" varchar(150) DEFAULT NULL,
    "token" char(36) NOT NULL DEFAULT uuid_generate_v4(),
    "gpu" varchar(20) DEFAULT NULL,
    "ram" float DEFAULT NULL,
    "cores" int DEFAULT NULL,
    "threads" int DEFAULT NULL,
    PRIMARY KEY ("id")
);

CREATE INDEX "hosts_name_IDX" on "hosts" ("name");

CREATE INDEX "hosts_ready_IDX" on "hosts" ("ready");

CREATE INDEX "hosts_updated_IDX" on "hosts" ("updated");

CREATE INDEX "hosts_url_IDX" on "hosts" ("job_url");

CREATE TABLE "host_gpus" (
    "id" char(36) DEFAULT uuid_generate_v4(),
    "host_id" char(36) DEFAULT NULL,
    "gpu_id" varchar(36) DEFAULT NULL,
    "slot_no" int DEFAULT 0,
    PRIMARY KEY ("id"),
    CONSTRAINT "host_gpus_ibfk_1" FOREIGN KEY ("host_id") REFERENCES "hosts" ("id"),
    CONSTRAINT "host_gpus_ibfk_2" FOREIGN KEY ("gpu_id") REFERENCES "gpu_models" ("id")
);

CREATE TABLE "host_services" (
    "host_id" varchar(36) NOT NULL,
    "service_id" varchar(36) NOT NULL,
    "gpu_id" char(36) DEFAULT NULL,
    PRIMARY KEY ("host_id"),
    CONSTRAINT "host_services_ibfk_1" FOREIGN KEY ("host_id") REFERENCES "hosts" ("id"),
    CONSTRAINT "host_services_ibfk_2" FOREIGN KEY ("service_id") REFERENCES "gpu_services" ("id")
);

CREATE INDEX "host_services_service_id_IDX" on "host_services" ("service_id");

CREATE INDEX "host_services_host_id_IDX" on "host_services" ("host_id");

CREATE TABLE "pending_requests" (
    "id" char(36) NOT NULL,
    "user_id" char(36) NOT NULL,
    "cmd" varchar(100) DEFAULT NULL,
    "container" varchar(100) DEFAULT NULL,
    "time_quota" float DEFAULT NULL,
    "inserted_at" timestamp DEFAULT NULL,
    PRIMARY KEY ("id")
);

CREATE INDEX "pending_requests_user_id_IDX" on "pending_requests" ("user_id");

CREATE INDEX "pending_requests_inserted_at_IDX" on "pending_requests" ("inserted_at");

CREATE INDEX "pending_requests_container_IDX" on "pending_requests" ("container");

CREATE TABLE "users" (
    "id" char(36) NOT NULL DEFAULT uuid_generate_v4(),
    "user" varchar(100) NOT NULL UNIQUE,
    "password" varchar(100) NOT NULL,
    "valid" boolean DEFAULT false,
    "expires" timestamp NOT NULL DEFAULT (current_timestamp + interval '1 year'),
    "paying" boolean DEFAULT NULL,
    PRIMARY KEY ("id")
);

CREATE INDEX "users_expires_IDX" on "users" ("expires");

CREATE TABLE "requests" (
    "id" char(36) NOT NULL DEFAULT uuid_generate_v4(),
    "user_id" char(36) DEFAULT NULL,
    "cmd" varchar(100) DEFAULT NULL,
    "time" float DEFAULT NULL,
    "executed_at" timestamp DEFAULT current_timestamp,
    PRIMARY KEY ("id"),
    CONSTRAINT "requests_ibfk_1" FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

CREATE INDEX "requests_user_id_IDX" on "requests" ("user_id");

CREATE TABLE "sessions" (
    "id" char(36) NOT NULL DEFAULT uuid_generate_v4(),
    "user_id" char(36) NOT NULL,
    "valid" boolean DEFAULT true,
    "expires" timestamp NOT NULL DEFAULT (current_timestamp + interval '1 day'),
    "time" float DEFAULT 0,
    PRIMARY KEY ("id"),
    CONSTRAINT "sessions_ibfk_1" FOREIGN KEY ("user_id") REFERENCES "users" ("id")
);

CREATE INDEX "sessions_user_id_IDX" on "sessions" ("user_id");

CREATE INDEX "sessions_valid_IDX" on "sessions" ("valid");

CREATE INDEX "sessions_expires_IDX" on "sessions" ("expires");

CREATE TABLE "disks" (
    "id" char(36) NOT NULL DEFAULT uuid_generate_v4(),
    "disk_no" int DEFAULT NULL,
    "type" varchar DEFAULT NULL,
    "capacity" float DEFAULT NULL,
    "read_speed" int DEFAULT NULL,
    "write_speed" int DEFAULT NULL,
    "host_id" char(36) DEFAULT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "disks_ibfk_1" FOREIGN KEY("host_id") REFERENCES "hosts" ("id")
);

CREATE TABLE "sensor_models" (
    "id" char(36) NOT NULL DEFAULT uuid_generate_v4(),
    "name" varchar DEFAULT NULL,
    "description" varchar DEFAULT NULL,
    "warning_temp" float DEFAULT NULL,
    "critical_temp" float DEFAULT NULL,
    "sensor_no" int DEFAULT NULL,
    PRIMARY KEY ("id")
);

CREATE INDEX "sensor_models_name_IDX" on "sensor_models" ("name");

CREATE TABLE "sensor_readings" (
    "id" char(36) NOT NULL DEFAULT uuid_generate_v4(),
    "sensor_model_id" char(36) DEFAULT NULL,
    "host_id" char(36) DEFAULT NULL,
    "reading" float DEFAULT NULL,
    "updated" timestamp DEFAULT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "sensor_readings_ibfk_1" FOREIGN KEY("host_id") REFERENCES "hosts" ("id"),
    CONSTRAINT "sensor_readings_ibfk_2" FOREIGN KEY("sensor_model_id") REFERENCES "hosts" ("id")
);

CREATE TABLE "roles" (
    "id" char(36) NOT NULL DEFAULT uuid_generate_v4(),
    "is_owner" boolean DEFAULT false,
    "is_renter" boolean DEFAULT false,
    "microservice_consumer" boolean DEFAULT false,
    "batch_consumer" boolean DEFAULT false,
    "user_id" char(36) NOT NULL,
    PRIMARY KEY ("id"),
    CONSTRAINT "roles_ibfk_1" FOREIGN KEY("user_id") REFERENCES "users" ("id")
);

CREATE INDEX "roles_is_owner_IDX" on "roles" ("is_owner");

CREATE INDEX "roles_is_renter_IDX" on "roles" ("is_renter");

CREATE INDEX "roles_microservice_consumer_IDX" on "roles" ("microservice_consumer");

CREATE INDEX "roles_batch_consumer_IDX" on "roles" ("batch_consumer");