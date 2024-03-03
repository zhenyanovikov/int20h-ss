import {
  Stack,
  FormControl,
  InputLabel,
  Select,
  MenuItem,
  TextField,
  Button,
  Box,
} from "@mui/material";
import { useFormik } from "formik";
import { useTranslation } from "react-i18next";
import { useInviteStudent } from "../../../api/student";
import { useGetGroups } from "../../../api/group";
import { getInviteStudentValidationSchema } from "../../../helpers/validation";

function InviteStudentForm() {
  const { t } = useTranslation();
  const { mutate: inviteStudentMutate, isLoading: isInviteTeacherLoading } =
    useInviteStudent();
  const { data: groupsData } = useGetGroups();

  const formik = useFormik({
    initialValues: {
      groupId: "",
      email: "",
    },
    validationSchema: getInviteStudentValidationSchema(),
    onSubmit: handleSubmit,
  });

  const groups = groupsData || [];

  return (
    <form onSubmit={formik.handleSubmit}>
      <Stack
        sx={{
          flexDirection: {
            sx: "column",
            md: "row",
          },
          gap: 2,
        }}
      >
        <Box sx={{ width: { sx: "100%", md: "184px" } }}>
          <FormControl size="small" fullWidth>
            <InputLabel id="groupId">
              {t("organisms.inviteStudentForm.form.groupId.label")}
            </InputLabel>
            <Select
              labelId="groupId"
              id="groupId"
              name="groupId"
              value={formik.values.groupId}
              label={t("organisms.inviteStudentForm.form.groupId.label")}
              onChange={formik.handleChange}
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
          id="email"
          name="email"
          size="small"
          label={t("organisms.inviteStudentForm.form.email.label")}
          value={formik.values.email}
          onChange={formik.handleChange}
          onBlur={formik.handleBlur}
          error={formik.touched.email && !!formik.errors.email}
          helperText={formik.touched.email && t(formik.errors.email)}
        />
        <Stack>
          <Button
            variant="contained"
            type="submit"
            disabled={isInviteTeacherLoading}
          >
            {t("organisms.inviteStudentForm.form.button.submit")}
          </Button>
        </Stack>
      </Stack>
    </form>
  );

  function handleSubmit(values, actions) {
    inviteStudentMutate(values);

    actions.resetForm();
  }
}

export default InviteStudentForm;
