# products-service
(Service backend Go - (Login-JWT-ORM)) Móviles P1

## Description
This repo is an complement of [articles_mobile](https://github.com/SantiagoGaonaC/articles_mobile)

## Features

- [X] LoginController
- [X] ValidationToken
- [X] Products
- [X] Favorites


  Each API functionality could be access as listed in the table below:

<table>
  <tbody>
    <tr>
      <th>Verb</th>
      <th>URI</th>
      <th>Auth Needed?</th>
      <th>Method</th>
      <th>Description</th>
    </tr>
    <tr>
      <td>POST</td>
      <td>/login</td>
      <td class='text-align:center'>No</td>
      <td>Authenticate/Login User</td>
      <td>Authenticate an user and return a JWT</td>
    </tr>
    <tr>
      <td>GET</td>
      <td>/validation-token</td>
      <td class='text-align:center'>No</td>
      <td>Validate Token</td>
      <td>Validate token/User (Bearer Authentication) and return a JWT</td>
    </tr>    
    <tr>
      <td>POST</td>
      <td>/products</td>
      <td class='text-align:center'>YES</td>
      <td>List Products</td>
      <td>Return all products</td>
    </tr>
    <tr>
      <td>POST</td>
      <td>/favorites/{id}</td>
      <td class='text-align:center'>YES</td>
      <td>Add Favorite</td>
      <td>Add favorite and return update all favorites</td>
    </tr>
    <tr>
      <td>DELETE</td>
      <td>/favorites/{id}</td>
      <td class='text-align:center'>YES</td>
      <td>Delete Favorite</td>
      <td>Delete favorite and return update all favorites</td>
    </tr>
  </tbody>
</table>

## Installation

### Clone repo
```bash
git clone https://github.com/SantiagoGaonaC/articles-service.git
```

### Run service 
```bash
docker-compose up -d
```
