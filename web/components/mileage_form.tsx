import { useEffect, useState } from "react";
import axios from "axios";
import CategorySelect from "./categorySelect";
import GrantSelect from "./grantSelect";
import { useAppContext } from "@/context/AppContext";
import { Axios_Credentials } from "@/types/users";
import ErrorDisplay from "./errorDisplay";
export default function MileageForm({
  new_request,
  request_id,
}: {
  new_request: boolean;
  request_id?: string;
}) {
  const { user_credentials } = useAppContext();
  const [mileageRequestInput, setMileageRequestInput] = useState({
    grant_id: "",
    date: new Date().toISOString(),
    category: "",
    starting_location: "",
    destination: "",
    trip_purpose: "",
    start_odometer: 0,
    end_odometer: 1,
    tolls: 0.0,
    parking: 0.0,
  });
  useEffect(() => {
    const fetchRequestInfo = async (request_id: string) => {
      const { data, status, statusText } = await axios.get(
        "/api/mileage/detail",
        {
          ...user_credentials,
          data: {
            mileage_id: request_id,
          },
        }
      );
      if (status != 200 || 201) {
        return (
          <ErrorDisplay
            message={statusText}
            path="GET /mileage/detail"
            error={data}
          />
        );
      }
      const {
        grant_id,
        date,
        category,
        starting_location,
        destination,
        trip_purpose,
        start_odometer,
        end_odometer,
        tolls,
        parking,
      } = data.data;
      setMileageRequestInput({
        grant_id,
        date,
        category,
        starting_location,
        destination,
        trip_purpose,
        start_odometer,
        end_odometer,
        tolls,
        parking,
      });
    };
    for (const field of Object.keys(mileageRequestInput)) {
      const input = document.getElementById(field) as
        | HTMLSelectElement
        | HTMLInputElement;
      if (
        input?.value != "" &&
        input?.value != null &&
        input?.value != undefined
      ) {
        document.getElementById(`invalid-${field}`)?.classList.add("hidden");
        document.getElementById(field)?.classList.remove("invalid-input");
      } else {
        document.getElementById(field)?.classList.add("invalid-input");
        document.getElementById(`invalid-${field}`)?.classList.remove("hidden");
      }
      if (field === "start_odometer" || field === "end_odometer") {
        if (
          mileageRequestInput.start_odometer > mileageRequestInput.end_odometer
        ) {
          document.getElementById(field)?.classList.add("invalid-input");
          document
            .getElementById("invalid-odometer-1")
            ?.classList.remove("hidden");
          document
            .getElementById("invalid-odometer-2")
            ?.classList.remove("hidden");
        } else {
          document.getElementById(field)?.classList.remove("invalid-input");
          document
            .getElementById("invalid-odometer-1")
            ?.classList.add("hidden");
          document
            .getElementById("invalid-odometer-2")
            ?.classList.add("hidden");
        }
      }
    }
    !new_request && request_id && fetchRequestInfo(request_id);
  }, [new_request, request_id, user_credentials, mileageRequestInput]);

  const handleChange = (e: any) => {
    e.preventDefault();
    const { name, value } = e.target;
    let new_state;
    switch (name.trim().toLowerCase()) {
      case "date":
        new_state = {
          ...mileageRequestInput,
          [name]: new Date(value).toISOString(),
        };
        break;
      case "start_odometer":
      case "end_odometer":
        new_state = {
          ...mileageRequestInput,
          [name]: parseInt(value),
        };
        break;
      case "tolls":
      case "parking":
        new_state = {
          ...mileageRequestInput,
          [name]: parseFloat(value),
        };
        break;
      default:
        new_state = {
          ...mileageRequestInput,
          [name]: value.trim().toLowerCase(),
        };
        break;
    }
    setMileageRequestInput(new_state);
  };
  const createMileage = async (config: Axios_Credentials) => {
    axios
      .post("/api/mileage", mileageRequestInput, config)
      .then((response) => console.log(response.data))
      .catch((error) => console.error(error));
  };
  const saveEdits = async (config: Axios_Credentials) => {
    const request_body = { ...mileageRequestInput, request_id };
    axios
      .put("/api/mileage", request_body, config)
      .then((res) => console.log(res))
      .catch((err) => console.error(err));
  };
  const handleSubmit = async (e: any) => {
    e.preventDefault();
    let res;
    if (new_request) {
      res = await createMileage(user_credentials);
    } else {
      res = await saveEdits(user_credentials);
    }
    console.log(res);
  };

  return (
    <form id="mileage-form">
      <CategorySelect
        state={mileageRequestInput}
        setState={setMileageRequestInput}
      />
      <span id="invalid-category" className="REJECTED field-span">
        <br />
        Category is Required
      </span>
      <GrantSelect
        state={mileageRequestInput}
        setState={setMileageRequestInput}
      />
      <span id="invalid-grant_id" className="REJECTED field-span">
        <br />
        Grant is Required
      </span>
      <h4>Trip Date</h4>
      <input
        type="datetime-local"
        name="date"
        id="date"
        onChange={handleChange}
        className="invalid-input"
      />
      <span id="invalid-date" className="REJECTED field-span">
        <br />
        Trip Date is Required
      </span>
      <h4>Starting Location</h4>
      <input
        name="starting_location"
        defaultValue={mileageRequestInput.starting_location}
        id="starting_location"
        className="invalid-input"
        maxLength={50}
        type="text"
        onChange={handleChange}
      />
      <br />
      <span id="invalid-starting_location" className="REJECTED field-span">
        Starting Location is Required
      </span>
      <span>{mileageRequestInput.starting_location.length}/50 characters</span>
      <h4>Destination</h4>
      <input
        name="destination"
        id="destination"
        className="invalid-input"
        defaultValue={mileageRequestInput.destination}
        maxLength={50}
        type="text"
        onChange={handleChange}
      />
      <br />
      <span id="invalid-destination" className="REJECTED field-span">
        Trip Destination is Required
      </span>
      <span>{mileageRequestInput.destination.length}/50 characters</span>
      <h4>Trip Purpose</h4>
      <textarea
        rows={5}
        maxLength={75}
        id="trip_purpose"
        name="trip_purpose"
        className="invalid-input"
        defaultValue={mileageRequestInput.trip_purpose}
        onChange={handleChange}
      />
      <br />
      <span id="invalid-trip_purpose" className="REJECTED field-span">
        Trip Purpose is Required
      </span>
      <span>{mileageRequestInput.trip_purpose.length}/75 characters</span>
      <h4>Start Odometer</h4>
      <input
        id="start_odometer"
        name="start_odometer"
        defaultValue={mileageRequestInput.start_odometer}
        type="number"
        className="invalid-input"
        onChange={handleChange}
      />
      <span id="invalid-odometer-1" className="REJECTED field-span">
        <br />
        Current Odometer Reading is Impossible
      </span>
      <span id="invalid-start_odometer" className="REJECTED field-span">
        <br />
        Start Odometer is Required
      </span>
      <h4>End Odometer</h4>
      <input
        id="end_odometer"
        name="end_odometer"
        defaultValue={mileageRequestInput.end_odometer}
        type="number"
        className="invalid-input"
        onChange={handleChange}
      />
      <span id="invalid-odometer-2" className="REJECTED field-span">
        <br />
        Current Odometer Reading is Impossible
      </span>
      <span id="invalid-end_odometer" className="REJECTED field-span">
        <br />
        End Odometer is Required
      </span>
      <h4>Tolls</h4>
      <input
        name="tolls"
        type="number"
        defaultValue={mileageRequestInput.tolls}
        onChange={handleChange}
      />
      <h4>Parking</h4>
      <input
        name="parking"
        type="number"
        defaultValue={mileageRequestInput.parking}
        onChange={handleChange}
      />
      <br />
      <br />
      <br />
      <a onClick={handleSubmit} className="archive-btn">
        Create Request
      </a>
      <br />
    </form>
  );
}
