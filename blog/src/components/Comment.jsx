import AccountCircleIcon from '@mui/icons-material/AccountCircle';

export default function Comment({ comment }) {
  console.log(comment);
  const removeTags = (str) => {
    return str.replace(/<\/?[^>]+(>|$)/g, '');
  };
  return (
    <>
      <div className="comment-wrapper">
        <div className="small-avatar">
          <AccountCircleIcon sx={{ fontSize: 70 }} />
        </div>
        <div className="comment-text-box">
          <div className="comment-author" ><strong>{comment.commentatorHandle}</strong></div>
          {/* <div className="comment-id">{comment.id}</div> */}
          <div className="comment-text">{removeTags(comment.text)}</div>
          <br />
        </div>
      </div>
    </>
  );
}
