<script lang="ts" setup>
interface Props {
  amount: number;
}

const formatter = new Intl.NumberFormat("en-US", {
  style: "currency",
  currency: "USD",
});

const props = defineProps<Props>();
const balance = computed(() => formatter.format(props.amount));
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
