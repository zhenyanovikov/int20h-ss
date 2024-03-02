// import { useGetStudents } from "../../../api/student";
import StudentsTemplate from "../../templates/StudentsTemplate/StudentsTemplate";

function StudentsPage() {
  // const { data: studentsData, isLoading: isGetStudentsLoading } =
  //   useGetFaculties();

  return (
    <StudentsTemplate
      // facultiesData={studentsData}
      // isGetFacultiesLoading={isGetStudentsLoading}
      studentsData={[]}
      isGetStudentsLoading={false}
    />
  );
}

export default StudentsPage;
