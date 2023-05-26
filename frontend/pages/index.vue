<script lang="ts" setup>
const headers = useRequestHeaders(["cookie"]);
const { data, error, pending } = useFetch("/api/me", { headers });

const isOperationModalShow = ref(false);
</script>

<template>
  <div>
    <div class="flex items-center justify-between mb-8">
      <strong class="text-4xl">Dashboard</strong>
      <button
        class="btn btn-primary btn-outline"
        @click="isOperationModalShow = true"
      >
        New operation
      </button>
    </div>
    <Balance :amount="data?.balance!" />

    <article class="card mb-8">
      <TableControl />
      <RecordTable />
    </article>
    <TablePagination />
    <OperationModal
      :show="isOperationModalShow"
      @close="isOperationModalShow = false"
    />
  </div>
</template>
