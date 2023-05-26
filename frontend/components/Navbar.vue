<template>
  <div class="bg-base-100 mb-8">
    <div class="container mx-auto">
      <div class="navbar">
        <div class="navbar-start">
          <strong class="text-2xl">Arithmetic</strong>
        </div>
        <div class="navbar-end">
          <button class="dropdown dropdown-end">
            <label tabindex="0" class="btn btn-ghost btn-circle avatar">
              <div class="w-10 rounded-full">
                <img
                  src="https://daisyui.com/images/stock/photo-1534528741775-53994a69daeb.jpg"
                />
              </div>
            </label>
            <ul
              tabindex="0"
              class="menu menu-compact dropdown-content mt-3 p-2 shadow bg-base-100 rounded-box w-52"
            >
              <li>
                <NuxtLink to="/api/logout" external>Logout</NuxtLink>
              </li>
            </ul>
          </button>
        </div>
      </div>
    </div>
  </div>
</template>
