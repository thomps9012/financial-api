import { useAppContext } from "@/context/AppContext";
import { Petty_Cash_Request } from "@/types/petty_cash";
import axios from "axios";
import { getCookie } from "cookies-next";
import { GetServerSidePropsContext } from "next";

function ProfilePettyCashPage({
  petty_cash_requests,
}: {
  petty_cash_requests: Petty_Cash_Request[];
}) {
  const { user_profile } = useAppContext();
  const { name } = user_profile;
  return (
    <main>
      <h1>Petty Cash Page for {name}</h1>
      <p>{JSON.stringify(petty_cash_requests, null, 2)}</p>
    </main>
  );
}

ProfilePettyCashPage.getInitialProps = async (
  ctx: GetServerSidePropsContext
) => {
  const credentials = getCookie("auth_credentials", {
    req: ctx.req,
    res: ctx.res,
  });
  if (credentials) {
    const { data } = await axios.get(
      "/me/petty_cash",
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

export default ProfilePettyCashPage;
