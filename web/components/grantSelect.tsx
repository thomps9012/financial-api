import { useAppContext } from "@/context/AppContext";
import { Grant } from "../types/grants";
export default function GrantSelect({
  state,
  setState,
}: {
  state: any;
  setState: any;
}) {
  const { grant_list } = useAppContext();
  let handleChange = (event: any) => {
    const { value } = event.target;
    const new_state = { ...state, grant_id: value.trim() };
    setState(new_state);
  };
  return (
    <>
      <h3>Grant</h3>
      <select name={state} onChange={handleChange} defaultValue={state}>
        <option value="" disabled hidden>
          Select Grant...
        </option>
        <option value="N/A">None</option>
        {grant_list.map((grant: Grant) => {
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
