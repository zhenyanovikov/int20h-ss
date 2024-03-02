import { useGetGroups } from "../../../api/group";
import GroupsTemplate from "../../templates/GroupsTemplate/GroupsTemplate";

function GroupsPage() {
  const { data: groupsData, isLoading: isGetGroupsLoading } = useGetGroups();

  return (
    <GroupsTemplate
      groupsData={groupsData}
      isGetGroupsLoading={isGetGroupsLoading}
    />
  );
}

export default GroupsPage;
