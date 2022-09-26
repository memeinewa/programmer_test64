# Exam 5

## Write sql statement to check user is login successfully?

### 1. Login successfully

### 2. Login failed

### SQL Answer
```
SELECT tb_Login.id, tb_emp.f_name, tb_emp.l_name, tb_position.desc_th FROM tb_Login
JOIN tb_emp ON tb_Login.id = tb_emp.id
JOIN tb_position ON tb_emp.pos_code = tb_position.pos_code
WHERE tb_Login.id = "?" and tb_Login.pass = "?"
```