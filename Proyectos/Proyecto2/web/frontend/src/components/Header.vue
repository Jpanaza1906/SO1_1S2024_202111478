<template>
  <header class="header">
    <div class="logo">
      <img src="../assets/logo.png" alt="Logo">
    </div>
    <!-- Enlace "Home" en la esquina derecha -->
    <nav class="nav">
      <button>
        <a href="#" class="btn2"><span class="spn2">HOME</span></a>
      </button>
    </nav>
  </header>
</template>

<script>
export default {
  name: 'HeaderComponent',
}
</script>

<style scoped>
.header {
  display: flex;
  justify-content: space-between;
  align-items: center;

  padding: 0.6rem;
  background-color: #ffffff77;
  margin: 0;
}

.logo img {
  max-height: 3rem;
  /* Establece el tamaño máximo del logo */
}

.nav {
  font-size: 1.2rem;
  cursor: pointer;

}

.btn2 {
  position: relative;
  display: inline-block;
  padding: 15px 30px;
  border: 2px solid #fefefe;
  text-transform: uppercase;
  color: #000;
  text-decoration: none;
  font-weight: 600;
  font-size: 20px;
  transition: 0.3s;
}

.btn2::before {
  content: '';
  position: absolute;
  top: -2px;
  left: -2px;
  width: calc(100% + 4px);
  height: calc(100% - -2px);
  background-color: #616161c4;
  transition: 0.3s ease-out;
  transform: scaleY(1);
}

.btn2::after {
  content: '';
  position: absolute;
  top: -2px;
  left: -2px;
  width: calc(100% + 4px);
  height: calc(100% - 50px);
  background-color: #616161c4;
  transition: 0.3s ease-out;
  transform: scaleY(1);
}

.btn2:hover::before {
  transform: translateY(-25px);
  height: 0;
}

.btn2:hover::after {
  transform: scaleX(0);
  transition-delay: 0.15s;
}

.btn2:hover {
  border: 2px solid #fefefe;
}

.btn2 span {
  position: relative;
  z-index: 3;
}

button {
  text-decoration: none;
  border: none;
  background-color: transparent;
}
</style>