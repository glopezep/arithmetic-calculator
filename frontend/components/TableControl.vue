<template>
  <div class="flex justify-between items-center">
    <input
      type="text"
      placeholder="Search"
      class="input input-bordered w-full max-w-xs mb-4"
    />
    <div class="flex space-x-4">
      <div class="form-control">
        <div class="input-group">
          <span>
            <i class="i-mdi-sort-ascending i-mdi-sort-descending"></i>
          </span>
          <select class="select select-bordered">
            <option disabled selected>Pick category</option>
            <option>T-shirts</option>
            <option>Mugs</option>
          </select>
        </div>
      </div>
    </div>
  </div>
</template>
