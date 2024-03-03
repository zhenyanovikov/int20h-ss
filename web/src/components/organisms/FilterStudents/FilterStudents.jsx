import {
  Stack,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  TextField,
  Box,
  Typography,
} from "@mui/material";
import { useTranslation } from "react-i18next";
import { useGetGroups } from "../../../api/group";

function FilterStudents({ name, onChangeName, groupId, onChangeGroupId }) {
  const { t } = useTranslation();

  const { data: groupsData } = useGetGroups();

  const groups = groupsData || [];

  return (
    <Stack spacing={2}>
      <Typography variant="h6" as="h2">
        {t("organisms.filterStudents.text.filter")}
      </Typography>
      <Stack
        sx={{
          flexDirection: {
            sx: "column",
            md: "row",
          },
          gap: 2,
        }}
      >
        <Box sx={{ minWidth: { sx: "100%", md: "184px" } }}>
          <FormControl fullWidth>
            <InputLabel id="facultyId">
              {t("organisms.filterStudents.filter.groupId.label")}
            </InputLabel>
            <Select
              labelId="groupId"
              value={groupId}
              label={t("organisms.filterStudents.filter.groupId.label")}
              onChange={handleChangeGroupId}
            >
              {groups.map((group) => (
                <MenuItem key={group.id} value={group.id}>
                  {group.name}
                </MenuItem>
              ))}
            </Select>
          </FormControl>
        </Box>
        <TextField
          fullWidth
          placeholder={t("organisms.filterStudents.filter.name.placeholder")}
          value={name}
          onChange={handleChangeName}
        />
      </Stack>
    </Stack>
  );

  function handleChangeGroupId(event) {
    onChangeGroupId(event.target.value);
  }

  function handleChangeName(event) {
    onChangeName(event.target.value);
  }
}

export default FilterStudents;
