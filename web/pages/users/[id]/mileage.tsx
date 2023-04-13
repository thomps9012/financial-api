import { GetServerSidePropsContext } from "next";

function UserMileagePage({ user_id }: { user_id: string }) {
  return (
    <main>
      <h1>Mileage page for {user_id}</h1>
    </main>
  );
}

UserMileagePage.getInitialProps = (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  return {
    user_id: id,
  };
};

export default UserMileagePage;
