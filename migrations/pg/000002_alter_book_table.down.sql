ALTER TABLE file_paths
    ADD book_id uuid,
    ADD CONSTRAINT fk_file_paths_book FOREIGN KEY(book_id) REFERENCES books(id);

ALTER TABLE books
    DROP CONSTRAINT fk_books_file_path,
    DROP COLUMN file_path_id;