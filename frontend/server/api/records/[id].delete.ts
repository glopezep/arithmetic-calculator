import { recordService } from "../../services/record-service";

export default defineEventHandler(async (event) => {
  const token = event.context.session.user?.token;
  const id = event.context.params?.id!;

  const res = await recordService.deleteRecord(id, {
    authorization: `Bearer ${token}`,
  });

  return res;
});
