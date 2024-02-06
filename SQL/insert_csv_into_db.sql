-- Table Structure
CREATE TABLE PERFORMANCE_LOG(
    idProyecto varchar(4),
    idPrueba varchar(2),
    idRegion varchar(3),
    idServidor int(2),
    timeStamp bigint,
    fecha datetime,
    elapsed varchar(10),
    label varchar(100),
    responseCode varchar(3),
    responseMessage varchar(20),
    threadName varchar(100),
    dataType varchar(20),
    success varchar(15),
    failureMessage varchar(50),
    bytes varchar(30),
    sentBytes varchar(30),
    grpThreads varchar(30),
    allThreads varchar(30),
    URL varchar(200),
    Latency varchar(30),
    IdleTime varchar(30),
    Connect varchar(30)
);

-- Insert CSV file into table
LOAD DATA INFILE '/path/to/file' INTO TABLE PERFORMANCE_LOG FIELDS TERMINATED BY '\t' LINES TERMINATED BY '\n' IGNORE 1 LINES;



-- Updating to easiest way
sudo mysql -u $user -p --local-infile=1 $database
LOAD DATA LOCAL INFILE '/var/lib/mysql-files/MyAccount.csv' INTO TABLE $table_name FIELDS TERMINATED BY ',' LINES TERMINATED BY '\n' IGNORE 1 LINES;
