import { Check_Request_Overview } from "@/types/check_requests";
import axios from "axios";
import { getCookie } from "cookies-next";
import { GetServerSidePropsContext } from "next";
import Link from "next/link";

function UserCheckPage({
  user_id,
  user_name,
  requests,
}: {
  user_id: string;
  user_name: string;
  requests: Check_Request_Overview[];
}) {
  return (
    <main>
      <h1>Check Requests for {user_name}</h1>
      <p>{JSON.stringify(requests, null, 2)}</p>
      <Link href={`/${user_id}`}>
        &larr; · &#8594; · &#x2192; Back to {user_name} Overview
      </Link>
    </main>
  );
}

UserCheckPage.getInitialProps = async (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  const credentials = getCookie("auth_credentials", {
    req: ctx.req,
    res: ctx.res,
  });
  if (!credentials) {
    return {
      user_id: "",
      user_name: "",
      requests: [],
    };
  }
  const auth = JSON.parse(credentials as string);
  const { data } = await axios.get("/user/check", {
    ...auth,
    data: {
      user_id: id,
    },
  });
  const { data: user_data } = await axios.get("/user/name", {
    ...auth,
    data: {
      user_id: id,
    },
  });
  return {
    user_id: id,
    user_name: user_data.data,
    requests: data.data,
  };
};

export default UserCheckPage;
