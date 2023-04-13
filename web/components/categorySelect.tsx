export default function CategorySelect({
  state,
  setState,
}: {
  state: any;
  setState: any;
}) {
  const Categories = [
    "IOP",
    "INTAKE",
    "PEERS",
    "ACT_TEAM",
    "IHBT",
    "PERKINS",
    "MENS_HOUSE",
    "NEXT_STEP",
    "LORAIN",
    "PREVENTION",
    "ADMINISTRATIVE",
    "FINANCE",
  ];
  const handleChange = (event: any) => {
    const { value } = event.target;
    const new_state = { ...state, category: value.trim().toUpperCase() };
    setState(new_state);
  };
  return (
    <>
      <h3>Request Category</h3>
      <select name={state} onChange={handleChange} defaultValue={state}>
        <option value="" disabled hidden>
          Select Category...
        </option>
        {Categories.map((category) => (
          <option value={category} key={category}>
            {category.toLowerCase().split("_").join(" ")}
          </option>
        ))}
      </select>
    </>
  );
}
