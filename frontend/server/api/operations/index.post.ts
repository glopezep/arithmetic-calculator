import { operationService } from "../../services/operation-service";

export default defineEventHandler(async (event) => {
  const token = event.context.session.user?.token;

  const body = await readBody(event);
  const res = await operationService.executeOperation(
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
