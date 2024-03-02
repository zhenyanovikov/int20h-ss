// import { useGetTeachers } from "../../../api/teacher";
import GroupsTemplate from "../../templates/GroupsTemplate/GroupsTemplate";

function GroupsPage() {
  // const { data: groupsData, isLoading: isGetGroupsLoading } =
  //   useGetTeachers();

  return (
    <GroupsTemplate
      // teachersData={groupsData}
      // isGetTeachersLoading={isGetGroupsLoading}
      groupsData={[]}
      isGetGroupsLoading={false}
    />
  );
}

export default GroupsPage;
