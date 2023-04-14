import { useAppContext } from "@/context/AppContext";
import { Mileage_Request } from "@/types/mileage";
import axios from "axios";
import { getCookie } from "cookies-next";
import { GetServerSidePropsContext } from "next";

function ProfileMileagePage({
  mileage_requests,
}: {
  mileage_requests: Mileage_Request[];
}) {
  const { user_profile } = useAppContext();
  const { name } = user_profile;
  return (
    <main>
      <h1>Mileage Page for {name}</h1>
      <p>{JSON.stringify(mileage_requests, null, 2)}</p>
    </main>
  );
}
ProfileMileagePage.getInitialProps = async (ctx: GetServerSidePropsContext) => {
  const credentials = getCookie("auth_credentials", {
    req: ctx.req,
    res: ctx.res,
  });
  if (credentials) {
    const { data } = await axios.get(
      "/me/mileage",
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
export default ProfileMileagePage;
