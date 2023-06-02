<script lang="ts" setup>
definePageMeta({
  middleware: ["auth"],
});

const filters = reactive({
  limit: 10,
  offset: 1,
  sort_by: "user_balance",
  order_by: "desc",
});

const headers = useRequestHeaders(["cookie"]);
const { data: userData, refresh: refreshUser } = useFetch("/api/me", {
  headers,
});
const { data: records, refresh: refreshRecords } = useFetch("/api/records", {
  headers,
  query: filters,
});

const isOperationModalShow = ref(false);

async function handleSave() {
  await Promise.all([refreshUser(), refreshRecords()]);

  isOperationModalShow.value = false;
}

async function handleDelete(id: string) {
  await $fetch(`/api/records/${id}`, {
    method: "DELETE",
  });

  await refreshRecords();
}
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
    <Balance :balance="userData?.balance!" />

    <article class="card mb-8">
      <TableControl
        v-model:sort-by="filters.sort_by"
        v-model:order-by="filters.order_by"
      />
      <RecordTable :records="records?.items" @delete="handleDelete" />
    </article>
    <TablePagination
      :has-next-page="records?.hasNextPage"
      v-model:limit="filters.limit"
      v-model:offset="filters.offset"
      @update:limit="(limit) => (filters.limit = limit)"
      @update:offset="(offset) => (filters.offset = offset)"
    />
    <OperationModal
      :show="isOperationModalShow"
      @close="isOperationModalShow = false"
      @save="handleSave"
    />
  </div>
</template>
