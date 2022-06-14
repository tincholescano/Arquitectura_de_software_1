import React, { useState } from "react";
import "./Home.css";
import logo from "./images/home.svg"
import cart from "./images/cart.svg"
import Cookies from "universal-cookie";

const Cookie = new Cookies();

async function getUserById(id){
    return await fetch('http://127.0.0.1:8080/user/' + id, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
}).then(response => response.json())

}

async function getCategories(){
  return await fetch('http://127.0.0.1:8080/categories', {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}

async function getProducts(){
  return await fetch('http://127.0.0.1:8080/products', {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}

async function getProductsByCategoryId(id){
  return await fetch('http://127.0.0.1:8080/products/' + id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}

async function getCategoryById(id){
  return await fetch('http://127.0.0.1:8080/category/' + id, {
    method: "GET",
    headers: {
      "Content-Type": "application/json"
    }
  }).then(response => response.json())
}

function goto(path){
  window.location = window.location.origin + path
}


function gotologin(){
  goto("/login")
}


function retry() {
  goto("/")
}

function productsByCategoryId(id, setter, categorySetter) {
  getProductsByCategoryId(id).then(response => {setter(response); Cookie.set("category", id); getCategoryById(id).then(category => categorySetter(category))})
}

function showCategories(categories, setter, categorySetter) {
  return categories.map((category, i) => <a onClick={() => productsByCategoryId(category.category_id, setter, categorySetter)} obj={category} key={category.category_id}>{category.name}</a>)
}

function addToCart(id, setCartItems){
  let cookie = Cookie.get("cart");

  if(cookie == undefined){
    Cookie.set("cart", id + ",1;", {path: "/"});
    setCartItems(1)
    return
  }
  let newCookie = ""
  let isNewItem = true
  let toCompare = cookie.split(";")
  let total = 0;
  toCompare.forEach((item) => {
    if(item != ""){
      let array = item.split(",")
      let item_id = array[0]
      let item_quantity = array[1]
      if(id == item_id){
        item_quantity = Number(item_quantity) + 1
        isNewItem = false
      }
      newCookie += item_id + "," + item_quantity + ";"
      total += Number(item_quantity);
    }
  });
  if(isNewItem){
    newCookie += id + ",1;"
    total += 1;
  }
  cookie = newCookie
  Cookie.set("cart", cookie, {path: "/"})
  Cookie.set("cartItems", total, {path: "/"})
  setCartItems(total)
  return
}

function showProducts(products, setCartItems){
  return products.map((product) =>
    <div class="col s2">
      <div class="product large" key={product.product_id} className="product">
        <div class="product-image">
        <img width="128px" height="300px" src={product.picture_url}  onError={(e) => (e.target.onerror = null, e.target.src = "./images/default.jpg")}/>
        </div>
        <div class="product-content">
          <span class="text-blue"><a className="name">{product.name}</a></span>
          <p>Cantidad disponible: {product.description}</p>
          <p>Precio: ${product.base_price}</p>
        </div>
        <div class="product-action">
          <a class="waves-effect waves-light btn yellow black-text" onClick={() => addToCart(product.product_id, setCartItems)}>Agregar al carrito</a>
        </div>
      </div>
    </div>
 )
}

function logout(){
  Cookie.set("user_id", -1, {path: "/"})
  document.location.reload()
}

function search(){
  let input, filter, a, i;
  input = document.getElementById("search");
  filter = input.value.toUpperCase();
  a = document.getElementsByClassName("product");
  for (i = 0; i < a.length; i++) {
    let txtValue = a[i].children[1].textContent || a[i].children[1].innerText;
    if (txtValue.toUpperCase().indexOf(filter) > -1) {
      a[i].style.display = "inherit";
    } else {
      a[i].style.display = "none";
    }
  }
  if(input.value.toUpperCase().length <= 0){
    for(i = 0; i < a.length; i++){
      a[i].style.display = "inherit";
    }
  }

}

function deleteCategory(){
  Cookie.set("category", 0, {path: "/"})
  goto("/")
}

function gotocart(){
  goto("/cart")
}


function Home() {
  const [isLogged, setIsLogged] = useState(false)
  const [user, setUser] = useState({})
  const [categories, setCategories] = useState([])
  const [products, setProducts] = useState([])
  const [category, setCategory] = useState("")
  const [cartItems, setCartItems] = useState("")


  if (Cookie.get("user_id") > -1 && !isLogged){
    getUserById(Cookie.get("user_id")).then(response => setUser(response))
    setIsLogged(true)
  }

  if (!(Cookie.get("user_id") > -1) && isLogged){
    setIsLogged(false)
  }
   
  if(categories){
    getCategories().then(response => setCategories(response))
  }

  if (products.length <= 0){
    getProducts().then(response => {setProducts(response)})
  }

  if (!cartItems && Cookie.get("cartItems")){
    setCartItems(Cookie.get("cartItems"))
  }

  const login = (
    <ul id="nav-mobile" class="right hide-on-med-and-down">
      <li><a onClick={gotocart} class="black-text">Carrito</a></li>
      <li><a onClick={logout} class="black-text">Cerrar Sesion</a></li>
    </ul>
  )

  return (
    <div className="home">
      
      <nav class=" yellow accent-2 ">
        <div class="nav-wrapper">
          <a href="/" class="brand-logo center blue-text text-darken-2">Tienda Libre</a>
          <ul id="nav-mobile" class="right hide-on-med-and-down">
            <li>{isLogged ? login : <a onClick={gotologin} class="black-text">Iniciar Sesion</a>}</li>
          </ul>
        </div>
        <div class="nav-content ">
          <form>
            <div class="input-field">
              <input class="yellow black-text" id="search" type="search" required onChange={search}/>
              <label class="label-icon" for="search"><i class="material-icons blue-text">Buscar</i></label>
              <i class="material-icons">X</i>
            </div>
          </form>
        </div>
      </nav>
      <br></br>
      <br></br>
      <div class="row" id="main">
        {Cookie.get("category") > 0 ? <a className="categoryFilter"> {category.name} <button className="delete" onClick={deleteCategory}>X</button> </a> : <a/>}
        {products.length > 0 ? showProducts(products, setCartItems) : <a> </a>}
        

      </div>

      <div id="mySidenav" className="sidenav">

        {categories > 0 ? showCategories(categories, setProducts, setCategory) : <a onClick={retry}> CARGANDO... </a>}
      </div>
    </div>
  );
}

export default Home;
