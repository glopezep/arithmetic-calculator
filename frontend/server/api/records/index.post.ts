import { recordService } from "../../services/record-service";

export default defineEventHandler(async (event) => {
  const token = event.context.session.user?.token;

  const body = await readBody(event);
  const res = await recordService.executeOperation(
    {
      id: body?.id,
      firstValue: parseInt(body?.firstValue),
      secondValue: parseInt(body?.secondValue),
    },
    {
      authorization: `Bearer ${token}`,
    }
  );

  return {
    data: res,
  };
});
