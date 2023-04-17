import { Mileage_Request } from "@/types/mileage";
import axios from "axios";
import { getCookie } from "cookies-next";
import { GetServerSidePropsContext } from "next";
import Link from "next/link";

function MileageRequestDetail({
  request_id,
  request,
}: {
  request: Mileage_Request;
  request_id: string;
}) {
  return (
    <main>
      <h1>Mileage Request Detail for {request_id}</h1>
      <Link href={`/mileage/edit/${request_id}`}>Edit Request</Link>
      <p>{JSON.stringify(request, null, 2)}</p>
    </main>
  );
}

MileageRequestDetail.getInitialProps = async (
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
  const { data } = await axios.get("/mileage/detail", {
    ...user_credentials,
    data: {
      mileage_id: id,
    },
  });
  return {
    request_id: id,
    request: data.data,
  };
};

export default MileageRequestDetail;
