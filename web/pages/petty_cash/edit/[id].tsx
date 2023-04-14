import PettyCashFrom from "@/components/petty_cash_form";
import { GetServerSidePropsContext } from "next";

function EditPettyCashRequest({ request_id }: { request_id: string }) {
  return (
    <main>
      <h1>Edit Petty Cash Request {request_id}</h1>
      <PettyCashFrom new_request={false} request_id={request_id} />
    </main>
  );
}
EditPettyCashRequest.getInitialProps = (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  return {
    request_id: id,
  };
};
export default EditPettyCashRequest;
