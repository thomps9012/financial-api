import Check_Request_Form from "@/components/check_request_form";
import { GetServerSidePropsContext } from "next";

function EditCheckRequest({ request_id }: { request_id: string }) {
  return (
    <main>
      <h1>Check Request Edit Page for {request_id}</h1>
      <div className="hr" />
      <Check_Request_Form new_request={false} request_id={request_id} />
    </main>
  );
}

EditCheckRequest.getInitialProps = (ctx: GetServerSidePropsContext) => {
  const { id } = ctx.query;
  return {
    request_id: id,
  };
};

export default EditCheckRequest;
