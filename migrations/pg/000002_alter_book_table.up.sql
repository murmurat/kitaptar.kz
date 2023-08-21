-- ALTER TABLE books
--     ADD file_path_id uuid,
--     ADD CONSTRAINT FOREIGN KEY(file_path_id) REFERENCES file_paths(id);
ALTER TABLE books
    ADD file_path_id uuid,
    ADD CONSTRAINT fk_books_file_path FOREIGN KEY(file_path_id) REFERENCES file_paths(id);

ALTER TABLE file_paths
    DROP CONSTRAINT fk_file_paths_book,
    DROP COLUMN book_id;

--  ALTER TABLE file_paths
--     DROP FOREIGN KEY `fk_file_paths_book`,
--     DROP COLUMN book_id;

