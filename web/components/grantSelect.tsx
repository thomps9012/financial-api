import { Grant } from "../types/grants";
export default function GrantSelect({
  state,
  setState,
}: {
  state: any;
  setState: any;
}) {
  const grants = [
    {
      id: "H79TI082369",
      name: "BCORR",
    },
    {
      id: "SOR_HOUSING",
      name: "SOR Recovery Housing",
    },
    {
      id: "SOR_PEER",
      name: "SOR Peer",
    },
    {
      id: "SOR_LORAIN",
      name: "SOR Lorain 2.0",
    },
    {
      id: "H79TI085495",
      name: "RAP AID (Recover from Addition to Prevent Aids)",
    },
    {
      id: "2020-JY-FX-0014",
      name: "JSBT (OJJDP) - Jump Start For A Better Tomorrow",
    },
    {
      id: "H79SP082264",
      name: "HIV Navigator",
    },
    {
      id: "H79SP082475",
      name: "SPF (HOPE 1)",
    },
    {
      id: "SOR_TWR",
      name: "SOR 2.0 - Together We Rise",
    },
    {
      id: "H79TI083370",
      name: "BSW (Bridge to Success Workforce)",
    },
    {
      id: "H79SM085150",
      name: "CCBHC",
    },
    {
      id: "H79TI083662",
      name: "IOP New Serenity Intensive Outpatient Program",
    },
    {
      id: "TANF",
      name: "TANF",
    },
    {
      id: "H79SP081048",
      name: "STOP Grant",
    },
    {
      id: "H79TI085410",
      name: "N MAT (NORA Medication-Assisted Treatment Program)",
    },
  ];
  let handleChange = (event: any) => {
    const { value } = event.target;
    const new_state = { ...state, grant_id: value.trim() };
    setState(new_state);
  };
  return (
    <>
      <h4>Grant</h4>
      <select
        id="grant_id"
        name={state}
        onChange={handleChange}
        defaultValue={state.grant_id}
      >
        <option value="" disabled hidden>
          Select Grant...
        </option>
        <option value="N/A">None</option>
        {grants.map((grant: Grant) => {
          const { id, name } = grant;
          return (
            <option key={id} value={id}>
              {name}
            </option>
          );
        })}
      </select>
    </>
  );
}
