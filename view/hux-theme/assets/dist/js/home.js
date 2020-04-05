$(document).ready(function () {
    initComment();
    $("time.timeago").timeago();
});

function initComment() {
	
    var $list = $('#comment-list');
    if (!$list.length) {
        return;
    }

    if (localStorage.getItem("comment-author")) {
        $('#author').val(localStorage.getItem("comment-author"));
        $('#email').val(localStorage.getItem("comment-email"));
        $('#website').val(localStorage.getItem("comment-website"));
        $('#avatar').attr("src", localStorage.getItem("comment-avatar"));
        $('.c-avatar').removeClass("null");
    }

    $('#comment-content').on("focus", function () {
        if ($('.c-avatar').hasClass("null")) {
            $('.c-avatar-field').remove();
            $('.c-info-fields').removeClass("hide");
        }
    });
    $('#comment-show').on("click", function () {
        $('#comment-show').hide();
        var form = $('#comment-form');
        form.removeClass("hide").appendTo($list);
    });
    $('#comment-cancel').on("click", function () {
        $('#comment-form').addClass("hide");
        $('#comment-show').show();
    });

    $('#comment-form').ajaxForm(function (json) {
        if (json.res) {
			console.log(json.comment);
            localStorage.setItem("comment-author", $('#author').val());
            localStorage.setItem("comment-email", $('#email').val());
            localStorage.setItem("comment-website", $('#website').val());
            localStorage.setItem("comment-avatar", json.comment.avatar);
            var tpl = $($('#comment-tpl').html());
            tpl.find(".comment-avatar").attr("src", json.comment.avatar).attr("alt", json.comment.avatar);
            tpl.find(".comment-name").attr("href", json.comment.website).text(json.comment.author);
            tpl.find(".comment-reply").attr("rel", json.comment.id);
            tpl.find(".comment-content").html("<p>" + json.comment.content + "</p>");
            tpl.find(".comment-message").html("留言正在审核...");
            tpl.attr("id", "comment-" + json.comment.id);
            if (json.comment.status == "approved") {
                tpl.find(".comment-check").remove();
            }
            var parentId = $('#comment-parent').val();
            if (parentId === '0') {
                $('#comment-list').append(tpl);
            } else {
                tpl.insertAfter($('#comment-' + parentId));
            }
            $('#comment-cancel').trigger("click");
            $('#comment-content').val("");
        } else {
            alert("评论失败,请检查信息是否填写正确");
        }
    });
    $list.on("click", ".comment-reply", function () {
        var id = $(this).attr("rel");
        var parentComment = $('#comment-' + id);
        var form = $('#comment-form');
        form.removeClass("hide").insertAfter($(this));
        $('#comment-show').hide();
        $('#comment-parent').val(id);
        $('.cancel-reply').show();
        var top = parentComment.offset().top;
        $('body,html').animate({scrollTop: top}, 500);
        return false;
    });
    $('.cancel-reply').on("click", function () {
        $('#comment-reply').empty();
        $('#comment-parent').val(0);
        $(this).hide();
        return false;
    });
}
