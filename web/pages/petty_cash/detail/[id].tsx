import { Petty_Cash_Request } from "@/types/petty_cash";
import axios from "axios";
import { getCookie } from "cookies-next";
import { GetServerSidePropsContext } from "next";
import Link from "next/link";

function PettyCashRequestDetails({
  request_id,
  request,
}: {
  request_id: string;
  request: Petty_Cash_Request;
}) {
  return (
    <main>
      <h1>Details for Petty Cash Request {request_id}</h1>
      <Link href={`/petty_cash/edit/${request_id}`}>Edit Request</Link>
      <p>{JSON.stringify(request, null, 2)}</p>
    </main>
  );
}

PettyCashRequestDetails.getInitialProps = async (
  ctx: GetServerSidePropsContext
) => {
  const { id } = ctx.query;
  const credentials = getCookie("auth_credentials", {
    req: ctx.req,
    res: ctx.res,
  });
  if (!credentials) {
    return {
      request_id: "",
      request: {},
    };
  }
  const user_credentials = JSON.parse(credentials as string);
  const { data } = await axios.get("/petty_cash/detail", {
    ...user_credentials,
    data: {
      petty_cash_id: id,
    },
  });
  return {
    request_id: id,
    request: data.data,
  };
};
export default PettyCashRequestDetails;
