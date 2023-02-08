interface IProps {
  image?: {
    file: string;
  };
  class?: string;
}
const ProfilePicture = (props: IProps) => {
  var profilePic = new Image();
  if (!props.image) {
    // If there is no image, return default image
    return (
      <svg
        className="h-10 w-10 text-gray-400 sm:h-12 sm:w-12"
        fill="currentColor"
        viewBox="0 0 20 20"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          fillRule="evenodd"
          d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z"
          clipRule="evenodd"
          transform="translate(-1.6, 0)"
        ></path>
      </svg>
    );
  }
  profilePic.src = props.image.file;
  if (document.getElementById("profilePicture")) {
    let img = document.getElementById("profilePicture") as HTMLImageElement;
    if (img !== null) {
      img.src = profilePic.src;
    }
  }

  return <img id="profilePicture" src="" alt="" className={props.class} />;
};

export default ProfilePicture;
