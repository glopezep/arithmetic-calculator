import { recordService } from "../../services/record-service";

export default defineEventHandler(async (event) => {
  const token = event.context.session.user?.token;
  const query = getQuery(event);

  const res = await recordService.listRecords(query, {
    authorization: `Bearer ${token}`,
  });

  return res;
});
