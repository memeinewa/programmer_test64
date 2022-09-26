# Exam 3

## Write sql statement to find people

### From Data
```
data.json
```

### Sql command
```
command.sql
```

```
SELECT population.sector_code, sector.sector_desc, 
  population.male, population.female FROM population 
LEFT JOIN sector ON population.sector_code = sector.sector_code;
```
