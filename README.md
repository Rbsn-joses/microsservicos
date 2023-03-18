Projetinho para testar conhecimentos com microsserviços

psql -h 'localhost' -U 'postgres' -d 'postgres'

curl -X POST http://localhost:3000/api/servidorweb/login/authentication -H 'Content-Type: application/json' -d'
{
   "username": "Monica Araujo",
   "email": "monica.araujo@gmail.com",
   "password": "testando456"
}'

curl  http://localhost:3000/api/servidorweb/cursos -H 'Content-Type: application/json' -H 'Cookie: jwt-token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NzkxNDc0MTEsInVzZXJuYW1lIjoiTW9uaWNhIEFyYXVqbyJ9.soeYRwvmMHN0ecPdoKKqZuAFWRBW7oiBs_nxplDkxYU;'
