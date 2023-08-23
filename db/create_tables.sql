CREATE TABLE striker (
    striker_name VARCHAR NOT NULL,
    strength INTEGER NOT NULL,
    speed INTEGER NOT NULL,
    shooting INTEGER NOT NULL,
    passing INTEGER NOT NULL,
    technique INTEGER NOT NULL,
    head VARCHAR NOT NULL,
    arms VARCHAR NOT NULL,
    body VARCHAR NOT NULL,
    legs VARCHAR NOT NULL,
    PRIMARY KEY (striker_name, head, arms, body, legs)
);