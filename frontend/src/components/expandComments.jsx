

const expandComments = () => {

    const handleAddComs = () => {
        return (
            <div>
                add Comments Here
            </div>
        )
    }
    return (
        <div id="commentBox">
            <img onclick={handleAddComs} id="plusIcon"src="src/images/plus.png"/>
            <div id="CommentContent">
                Comment Content
            </div>
        </div>
    )
}

export default expandComments;