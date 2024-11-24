CREATE TABLE
    IF NOT EXISTS autos (
        id serial primary key not null unique,
        brand varchar(255) not null,
        model varchar(255) not null,
        statenumber varchar(9) not null unique
    );

CREATE INDEX IF NOT EXISTS idx_state_number ON autos (statenumber);

CREATE TABLE
    IF NOT EXISTS contragents (
        id serial primary key not null unique,
        name varchar(255) not null,
        address varchar(255) not null,
        innkpp varchar(20) not null unique
    );

CREATE INDEX IF NOT EXISTS idx_inn_kpp ON contragents (innkpp);

CREATE TABLE
    IF NOT EXISTS dispetchers (
        id serial primary key not null unique,
        fullname varchar(255) not null
    );

CREATE INDEX IF NOT EXISTS idx ON dispetchers (id);

CREATE TABLE
    IF NOT EXISTS drivers (
        id serial primary key not null unique,
        fullname varchar(255) not null,
        license varchar(10) not null unique,
        class varchar(255) not null
    );

CREATE INDEX IF NOT EXISTS idx_license ON drivers (license);

CREATE TABLE
    IF NOT EXISTS mehanics (
        id serial primary key not null unique,
        fullname varchar(255) not null
    );

CREATE INDEX IF NOT EXISTS idx ON mehanics (id);

CREATE TABLE
    IF NOT EXISTS organizations (
        id serial primary key not null unique,
        name varchar(255) not null,
        address varchar(255) not null,
        chief varchar(255) not null,
        financialchief varchar(255) not null,
        innkpp varchar(20) not null unique
    );

CREATE INDEX IF NOT EXISTS idx_inn_kpp ON organizations (innkpp);

CREATE TABLE
    IF NOT EXISTS bankaccounts (
        id serial primary key not null unique,
        bankaccountnumber varchar(20) not null unique,
        bankname varchar(255) not null,
        bankidnumber varchar(9) not null,
        organizationid int references organizations (id) on delete cascade on update cascade not null
    );

CREATE INDEX IF NOT EXISTS idx_bank_account_number ON bankaccounts (bankaccountnumber);

CREATE TABLE
    IF NOT EXISTS putlistheaders (
        id serial primary key not null unique,
        userid int not null,
        number int not null unique,
        bankaccountid int references bankaccounts (id) on delete cascade on update cascade not null,
        datewith date not null,
        datefor date not null,
        autoid int references autos (id) on delete cascade on update cascade not null,
        driverid int references drivers (id) on delete cascade on update cascade not null,
        dispetcherid int references dispetchers (id) on delete cascade on update cascade not null,
        mehanicid int references mehanics (id) on delete cascade on update cascade not null
    );

CREATE INDEX IF NOT EXISTS idx_number ON putlistheaders (number);

CREATE TABLE
    IF NOT EXISTS putlistbodies (
        id serial primary key not null unique,
        putlistheadernumber int references putlistheaders (number) on delete cascade on update cascade not null,
        number int not null,
        contragentid int references contragents (id) on delete cascade on update cascade not null,
        item varchar(255) not null,
        timewith timestamp not null,
        timefor timestamp not null
    );

CREATE INDEX IF NOT EXISTS idx ON putlistbodies (id);

INSERT INTO
    autos (brand, model, statenumber)
VALUES
    ('Volvo', 'VMF', 'В576ЛО123');

INSERT INTO
    contragents (name, address, innkpp)
VALUES
    (
        'ООО Тест',
        'ул. Тестовая 65',
        '123456789/1234567890'
    );

INSERT INTO
    dispetchers (fullname)
VALUES
    ('Иванов Иван Иванович');

INSERT INTO
    mehanics (fullname)
VALUES
    ('Петров Петр Петрович');

INSERT INTO
    drivers (fullname, license, class)
VALUES
    (
        'Николай Николаев Николаевич',
        '1234567890',
        'A, B, C'
    );

INSERT INTO
    organizations (name, address, chief, financialchief, innkpp)
VALUES
    (
        'ООО Пример',
        'ул. Примера 65',
        'Деканоидзе Давид Вахтангович',
        'Деканоидзе Даниэл Вахтангович',
        '123456789/1234567890'
    );

INSERT INTO
    bankaccounts (
        bankaccountnumber,
        bankname,
        bankidnumber,
        organizationid
    )
VALUES
    ('12345678900987654321', 'Т-Банк', '123456789', 1);

INSERT INTO
    putlistheaders (
        userid,
        number,
        bankaccountid,
        datewith,
        datefor,
        autoid,
        driverid,
        dispetcherid,
        mehanicid
    )
VALUES
    (1, 1, 1, '2024-11-13', '2024-11-14', 1, 1, 1, 1);

INSERT INTO
    putlistbodies (
        putlistheadernumber,
        number,
        contragentid,
        item,
        timewith,
        timefor
    )
VALUES
    (
        1,
        1,
        1,
        'Стул',
        '2024-11-13 15:00:00',
        '2024-11-14 17:00:00'
    );