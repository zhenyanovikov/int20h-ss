import { useGetTeachers } from "../../../api/teacher";
import TeachersTemplate from "../../templates/TeachersTemplate/TeachersTemplate";

function TeachersPage() {
  const { data: teachersData, isLoading: isGetTeachersLoading } =
    useGetTeachers();

  return (
    <TeachersTemplate
      teachersData={teachersData}
      isGetTeachersLoading={isGetTeachersLoading}
    />
  );
}

export default TeachersPage;
