import { GetServerSidePropsContext } from "next";

function UserOverviewPage({ user_id }: { user_id: string }) {
  return (
    <main>
      <h1>Overview page for {user_id}</h1>
    </main>
  );
}

UserOverviewPage.getInitialProps = (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  return {
    user_id: id,
  };
};

export default UserOverviewPage;
