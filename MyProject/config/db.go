package config

type Database struct{
    DbName string
    DbUser string
    DbPass string
}

func GetDbconfig() Database{
    return Database{
        DbName : "employee",
        DbUser : "root",
        DbPass : "90147",
    }
}