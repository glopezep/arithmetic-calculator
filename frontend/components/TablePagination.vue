<script lang="ts" setup>
defineProps(["limit", "offset", "hasNextPage"]);
defineEmits(["update:limit", "update:offset"]);
</script>

<template>
  <div class="btn-group flex items-center justify-end">
    <button
      class="btn"
      @click="$emit('update:offset', offset === 1 ? 1 : offset - 1)"
    >
      «
    </button>
    <button class="btn">Page {{ offset }}</button>
    <button
      class="btn"
      @click="$emit('update:offset', hasNextPage ? offset + 1 : offset)"
    >
      »
    </button>
    <input
      type="number"
      class="input input-bordered w-32"
      :value="limit"
      @input="$emit('update:limit', $event.target.value)"
    />
  </div>
</template>
