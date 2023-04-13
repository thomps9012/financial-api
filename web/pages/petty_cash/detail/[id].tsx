import { GetServerSidePropsContext } from "next";

function PettyCashRequestDetails({ request_id }: { request_id: string }) {
  return (
    <main>
      <h1>Details for Petty Cash Request {request_id}</h1>
    </main>
  );
}

PettyCashRequestDetails.getInitialProps = (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  return {
    request_id: id,
  };
};
export default PettyCashRequestDetails;
