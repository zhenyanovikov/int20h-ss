// import { useGetTeachers } from "../../../api/teacher";
import TeachersTemplate from "../../templates/TeachersTemplate/TeachersTemplate";

function TeachersPage() {
  // const { data: teachersData, isLoading: isGetTeachersLoading } =
  //   useGetTeachers();

  return (
    <TeachersTemplate
      // teachersData={studentsData}
      // isGetTeachersLoading={isGetTeachersLoading}
      teachersData={[]}
      isGetTeachersLoading={false}
    />
  );
}

export default TeachersPage;
