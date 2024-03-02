import { useGetFaculties } from "../../../api/faculty";
import FacultiesTemplate from "../../templates/FacultiesTemplate/FacultiesTemplate";

function FacultiesPage() {
  const { data: facultiesData, isLoading: isGetFacultiesLoading } =
    useGetFaculties();

  return (
    <FacultiesTemplate
      facultiesData={facultiesData}
      isGetFacultiesLoading={isGetFacultiesLoading}
    />
  );
}

export default FacultiesPage;
