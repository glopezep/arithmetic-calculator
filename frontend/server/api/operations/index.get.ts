import { operationService } from "../../services/operation-service";

export default defineEventHandler(async (event) => {
  const token = event.context.session.user?.token;

  const res = await operationService.listOperations({
    authorization: `Bearer ${token}`,
  });

  return res;
});
