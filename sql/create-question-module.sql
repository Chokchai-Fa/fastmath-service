CREATE TABLE "ListOperation" (
  "list_operation_code" varchar PRIMARY KEY,
  "list_operation" jsonb,
  "diffculty_code" varchar
);

CREATE TABLE "Diffculty" (
  "diffculty_code" varchar PRIMARY KEY,
  "diffculty_description" varchar
);

CREATE TABLE "Question" (
  "question_id" int PRIMARY KEY,
  "question_flag" varchar,
  "min_list_operation_code" varchar,
  "question" varchar,
  "question_result" varchar
);

CREATE TABLE "Solution" (
  "solution_id" int PRIMARY KEY,
  "question_id" int,
  "solution_result" varchar,
  "solution_score" int,
  "list_operation_code" varchar,
  "solution" varchar
);

ALTER TABLE "ListOperation" ADD FOREIGN KEY ("diffculty_code") REFERENCES "Diffculty" ("diffculty_code");

ALTER TABLE "Question" ADD FOREIGN KEY ("min_list_operation_code") REFERENCES "ListOperation" ("list_operation_code");

ALTER TABLE "Solution" ADD FOREIGN KEY ("question_id") REFERENCES "Question" ("question_id");

ALTER TABLE "Solution" ADD FOREIGN KEY ("list_operation_code") REFERENCES "ListOperation" ("list_operation_code");
