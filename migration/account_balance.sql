CREATE TABLE report (
    id SERIAL PRIMARY KEY,
    file_name TEXT NOT NULL,
    date TIMESTAMP NOT NULL,
    isProcessed BOOLEAN NOT NULL,
    total FLOAT,
    avrDebit FLOAT,
    avrCredit FLOAT,
    send_to TEXT[],
    error TEXT,
    CONSTRAINT unique_file_date UNIQUE (file_name, date)
);

CREATE TABLE transactions_by_month (
    id SERIAL PRIMARY KEY,
    month INTEGER NOT NULL,
    num_transactions INTEGER NOT NULL,
    report_id INTEGER NOT NULL,
    CONSTRAINT unique_tr_by_month UNIQUE (report_id, month),
    CONSTRAINT fk_report_id FOREIGN KEY (report_id) REFERENCES report(id)
);

CREATE TABLE transactions (
    id SERIAL PRIMARY KEY,
    report_id INTEGER REFERENCES report(id),
    number INTEGER NOT NULL,
    date DATE NOT NULL,
    movement VARCHAR(20) NOT NULL,
    value FLOAT NOT NULL
);

CREATE TABLE ignored_transactions (
    id SERIAL PRIMARY KEY,
    report_id INTEGER REFERENCES report(id),
    ignored_id TEXT NOT NULL,
    date TEXT NOT NULL,
    transaction TEXT NOT NULL,
    reason TEXT NOT NULL
);
