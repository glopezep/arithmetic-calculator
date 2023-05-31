<script lang="ts" setup>
interface Props {
  balance: number;
}

const props = defineProps<Props>();
</script>

<template>
  <article class="card card-bordered mb-16 w-1/2 bg-base-100">
    <div class="card-body">
      <div class="stat">
        <div class="stat-title text-xl">Balance</div>
        <div class="stat-value">{{ balance }}</div>
      </div>
    </div>
  </article>
</template>
