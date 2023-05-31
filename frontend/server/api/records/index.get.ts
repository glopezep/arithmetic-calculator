import { recordService } from "../../services/record-service";

export default defineEventHandler(async (event) => {
  const token = event.context.session.user?.token;

  const res = await recordService.listRecords({
    authorization: `Bearer ${token}`,
  });

  return res;
});
