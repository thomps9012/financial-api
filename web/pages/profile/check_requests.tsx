import { useAppContext } from "@/context/AppContext";
import { Check_Request } from "@/types/check_requests";
import axios from "axios";
import { getCookie } from "cookies-next";
import { GetServerSidePropsContext } from "next";

function ProfileCheckPage({
  check_requests,
}: {
  check_requests: Check_Request[];
}) {
  const { user_profile } = useAppContext();
  const { name } = user_profile;
  return (
    <main>
      <h1>Check Request Page for {name}</h1>
      <p>{JSON.stringify(check_requests, null, 2)}</p>
    </main>
  );
}
ProfileCheckPage.getInitialProps = async (ctx: GetServerSidePropsContext) => {
  const credentials = getCookie("auth_credentials", {
    req: ctx.req,
    res: ctx.res,
  });
  if (credentials) {
    const { data } = await axios.get(
      "/me/check",
      JSON.parse(credentials as string)
    );
    return {
      petty_cash_requests: data.data,
    };
  } else {
    return {
      petty_cash_requests: [],
    };
  }
};
export default ProfileCheckPage;
